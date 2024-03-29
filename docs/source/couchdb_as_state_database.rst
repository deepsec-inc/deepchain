CouchDB as the State Database
=============================

State Database options
----------------------

State database options include LevelDB and CouchDB. LevelDB is the default key-value state
database embedded in the peer process. CouchDB is an optional alternative external state database.
Like the LevelDB key-value store, CouchDB can store any binary data that is modeled in chaincode
(CouchDB attachment functionality is used internally for non-JSON binary data). But as a JSON
document store, CouchDB additionally enables rich query against the chaincode data, when chaincode
values (e.g. assets) are modeled as JSON data.

Both LevelDB and CouchDB support core chaincode operations such as getting and setting a key
(asset), and querying based on keys. Keys can be queried by range, and composite keys can be
modeled to enable equivalence queries against multiple parameters. For example a composite
key of ``owner,asset_id`` can be used to query all assets owned by a certain entity. These key-based
queries can be used for read-only queries against the ledger, as well as in transactions that
update the ledger.

If you model assets as JSON and use CouchDB, you can also perform complex rich queries against the
chaincode data values, using the CouchDB JSON query language within chaincode. These types of
queries are excellent for understanding what is on the ledger. Proposal responses for these types
of queries are typically useful to the client application, but are not typically submitted as
transactions to the ordering service. In fact, there is no guarantee the result set is stable
between chaincode execution and commit time for rich queries, and therefore rich queries
are not appropriate for use in update transactions, unless your application can guarantee the
result set is stable between chaincode execution time and commit time, or can handle potential
changes in subsequent transactions. For example, if you perform a rich query for all assets
owned by Alice and transfer them to Bob, a new asset may be assigned to Alice by another
transaction between chaincode execution time and commit time, and you would miss this "phantom"
item.

CouchDB runs as a separate database process alongside the peer, therefore there are additional
considerations in terms of setup, management, and operations. You may consider starting with the
default embedded LevelDB, and move to CouchDB if you require the additional complex rich queries.
It is a good practice to model chaincode asset data as JSON, so that you have the option to perform
complex rich queries if needed in the future.

.. note:: The key for a CouchDB JSON document cannot begin with an underscore ("_").  Also, a JSON
   document cannot use the following field names at the top level.  These are reserved for internal use.

   - ``Any field beginning with an underscore, "_"``
   - ``~version``

Using CouchDB from Chaincode
----------------------------

Chaincode queries
~~~~~~~~~~~~~~~~~

Most of the `chaincode shim APIs <https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim#ChaincodeStubInterface>`__
can be utilized with either LevelDB or CouchDB state database, e.g. ``GetState``, ``PutState``,
``GetStateByRange``, ``GetStateByPartialCompositeKey``. Additionally when you utilize CouchDB as
the state database and model assets as JSON in chaincode, you can perform rich queries against
the JSON in the state database by using the ``GetQueryResult`` API and passing a CouchDB query string.
The query string follows the `CouchDB JSON query syntax <http://docs.couchdb.org/en/2.1.1/api/database/find.html>`__.

The `marbles02 fabric sample <https://github.com/hyperledger/fabric-samples/blob/master/chaincode/marbles02/go/marbles_chaincode.go>`__
demonstrates use of CouchDB queries from chaincode. It includes a ``queryMarblesByOwner()`` function
that demonstrates parameterized queries by passing an owner id into chaincode. It then queries the
state data for JSON documents matching the docType of “marble” and the owner id using the JSON query
syntax:

.. code:: bash

  {"selector":{"docType":"marble","owner":<OWNER_ID>}}

.. couchdb-pagination:

CouchDB pagination
^^^^^^^^^^^^^^^^^^

Fabric supports paging of query results for rich queries and range based queries.
APIs supporting pagination allow the use of page size and bookmarks to be used for
both range and rich queries. To support efficient pagination, the Fabric
pagination APIs must be used. Specifically, the CouchDB ``limit`` keyword will
not be honored in CouchDB queries since Fabric itself manages the pagination of
query results and implicitly sets the pageSize limit that is passed to CouchDB.

If a pageSize is specified using the paginated query APIs (``GetStateByRangeWithPagination()``,
``GetStateByPartialCompositeKeyWithPagination()``, and ``GetQueryResultWithPagination()``),
a set of results (bound by the pageSize) will be returned to the chaincode along with
a bookmark. The bookmark can be returned from chaincode to invoking clients,
which can use the bookmark in a follow on query to receive the next "page" of results.

The pagination APIs are for use in read-only transactions only, the query results
are intended to support client paging requirements. For transactions
that need to read and write, use the non-paginated chaincode query APIs. Within
chaincode you can iterate through result sets to your desired depth.

Regardless of whether the pagination APIs are utilized, all chaincode queries are
bound by ``totalQueryLimit`` (default 100000) from ``core.yaml``. This is the maximum
number of results that chaincode will iterate through and return to the client,
in order to avoid accidental or malicious long-running queries.

.. note:: Regardless of whether chaincode uses paginated queries or not, the peer will
          query CouchDB in batches based on ``internalQueryLimit`` (default 1000)
          from ``core.yaml``. This behavior ensures reasonably sized result sets are
          passed between the peer and CouchDB when executing chaincode, and is
          transparent to chaincode and the calling client.

An example using pagination is included in the :doc:`couchdb_tutorial` tutorial.

CouchDB indexes
~~~~~~~~~~~~~~~

Indexes in CouchDB are required in order to make JSON queries efficient and are required for
any JSON query with a sort. Indexes can be packaged alongside chaincode in a
``/META-INF/statedb/couchdb/indexes`` directory. Each index must be defined in its own
text file with extension ``*.json`` with the index definition formatted in JSON following the
`CouchDB index JSON syntax <http://docs.couchdb.org/en/2.1.1/api/database/find.html#db-index>`__.
For example, to support the above marble query, a sample index on the ``docType`` and ``owner``
fields is provided:

.. code:: bash

  {"index":{"fields":["docType","owner"]},"ddoc":"indexOwnerDoc", "name":"indexOwner","type":"json"}

The sample index can be found `here <https://github.com/hyperledger/fabric-samples/blob/master/chaincode/marbles02/go/META-INF/statedb/couchdb/indexes/indexOwner.json>`__.

Any index in the chaincode’s ``META-INF/statedb/couchdb/indexes`` directory
will be packaged up with the chaincode for deployment. When the chaincode is
both installed on a peer and instantiated on one of the peer’s channels, the
index will automatically be deployed to the peer’s channel and chaincode
specific state database (if it has been configured to use CouchDB). If you
install the chaincode first and then instantiate the chaincode on the channel,
the index will be deployed at chaincode **instantiation** time. If the
chaincode is already instantiated on a channel and you later install the
chaincode on a peer, the index will be deployed at chaincode **installation**
time.

Upon deployment, the index will automatically be utilized by chaincode queries. CouchDB can automatically
determine which index to use based on the fields being used in a query. Alternatively, in the
selector query the index can be specified using the ``use_index`` keyword.

The same index may exist in subsequent versions of the chaincode that gets installed. To change the
index, use the same index name but alter the index definition. Upon installation/instantiation, the index
definition will get re-deployed to the peer’s state database.

If you have a large volume of data already, and later install the chaincode, the index creation upon
installation may take some time. Similarly, if you have a large volume of data already and instantiate
a subsequent version of the chaincode, the index creation may take some time. Avoid calling chaincode
functions that query the state database at these times as the chaincode query may time out while the
index is getting initialized. During transaction processing, the indexes will automatically get refreshed
as blocks are committed to the ledger.

CouchDB Configuration
---------------------

CouchDB is enabled as the state database by changing the ``stateDatabase`` configuration option from
goleveldb to CouchDB. Additionally, the ``couchDBAddress`` needs to configured to point to the
CouchDB to be used by the peer. The username and password properties should be populated with
an admin username and password if CouchDB is configured with a username and password. Additional
options are provided in the ``couchDBConfig`` section and are documented in place. Changes to the
*core.yaml* will be effective immediately after restarting the peer.

You can also pass in docker environment variables to override core.yaml values, for example
``CORE_LEDGER_STATE_STATEDATABASE`` and ``CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS``.

Below is the ``stateDatabase`` section from *core.yaml*:

.. code:: bash

    state:
      # stateDatabase - options are "goleveldb", "CouchDB"
      # goleveldb - default state database stored in goleveldb.
      # CouchDB - store state database in CouchDB
      stateDatabase: goleveldb
      # Limit on the number of records to return per query
      totalQueryLimit: 10000
      couchDBConfig:
         # It is recommended to run CouchDB on the same server as the peer, and
         # not map the CouchDB container port to a server port in docker-compose.
         # Otherwise proper security must be provided on the connection between
         # CouchDB client (on the peer) and server.
         couchDBAddress: couchdb:5984
         # This username must have read and write authority on CouchDB
         username:
         # The password is recommended to pass as an environment variable
         # during start up (e.g. LEDGER_COUCHDBCONFIG_PASSWORD).
         # If it is stored here, the file must be access control protected
         # to prevent unintended users from discovering the password.
         password:
         # Number of retries for CouchDB errors
         maxRetries: 3
         # Number of retries for CouchDB errors during peer startup
         maxRetriesOnStartup: 10
         # CouchDB request timeout (unit: duration, e.g. 20s)
         requestTimeout: 35s
         # Limit on the number of records per each CouchDB query
         # Note that chaincode queries are only bound by totalQueryLimit.
         # Internally the chaincode may execute multiple CouchDB queries,
         # each of size internalQueryLimit.
         internalQueryLimit: 1000
         # Limit on the number of records per CouchDB bulk update batch
         maxBatchUpdateSize: 1000
         # Warm indexes after every N blocks.
         # This option warms any indexes that have been
         # deployed to CouchDB after every N blocks.
         # A value of 1 will warm indexes after every block commit,
         # to ensure fast selector queries.
         # Increasing the value may improve write efficiency of peer and CouchDB,
         # but may degrade query response time.
         warmIndexesAfterNBlocks: 1

CouchDB hosted in docker containers supplied with Hyperledger Fabric have the
capability of setting the CouchDB username and password with environment
variables passed in with the ``COUCHDB_USER`` and ``COUCHDB_PASSWORD`` environment
variables using Docker Compose scripting.

For CouchDB installations outside of the docker images supplied with Fabric,
the
`local.ini file of that installation
<http://docs.couchdb.org/en/2.1.1/config/intro.html#configuration-files>`__
must be edited to set the admin username and password.

Docker compose scripts only set the username and password at the creation of
the container. The *local.ini* file must be edited if the username or password
is to be changed after creation of the container.

.. note:: CouchDB peer options are read on each peer startup.

Good practices for queries
--------------------------

Avoid using chaincode for queries that will result in a scan of the entire
CouchDB database. Full length database scans will result in long response
times and will degrade the performance of your network. You can take some of
the following steps to avoid long queries:

- When using JSON queries:

    * Be sure to create indexes in the chaincode package.
    * Avoid query operators such as ``$or``, ``$in`` and ``$regex``, which lead
      to full database scans.

- For range queries, composite key queries, and JSON queries:

    * Utilize paging support (as of v1.3) instead of one large result set.

- If you want to build a dashboard or collect aggregate data as part of your
  application, you can query an off-chain database that replicates the data
  from your blockchain network. This will allow you to query and analyze the
  blockchain data in a data store optimized for your needs, without degrading
  the performance of your network or disrupting transactions. To achieve this,
  applications may use block or chaincode events to write transaction data
  to an off-chain database or analytics engine. For each block received, the block
  listener application would iterate through the block transactions and build a
  data store using the key/value writes from each valid transaction's ``rwset``.
  The :doc:`peer_event_services` provide replayable events to ensure the
  integrity of downstream data stores.

.. Licensed under Creative Commons Attribution 4.0 International License
   https://creativecommons.org/licenses/by/4.0/
