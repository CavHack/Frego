package main

import "flag"

var (

    Parser                          = flag.String("parser", "dimacs", "Parser type. Available Parsers: dimacs")
    VerticesFilePath                = flag.String("vertices.filePath", "", "Vertices input file path")
    EdgesFilePath                   = flag.String("edges.filePath", "", "Edges input file path")
    ReadBatchSize                   = flag.Int("read.batch.size", 10000, "Insert batch size")
    Store                           = flag.String("store", "Cassandra", "Store type. Available stores: Cassandra")
    CassandraHosts                  = flag.String("cassandra.hosts", "127.0.100.1, 127.0.100.2", "Cassandra hosts separated by ','")
    CassandraKeyspace               = flag.String("cassandra.keyspace", "frego", "cassandra keyspace")
    CassandraReplicationFactor      = flag.Int("cassandra.replication_factor", 3, "Cassandra keyspace")
    CassandraVerticesTable          = flag.String("cassandra.vertices_table", "", "Cassandra table name")
    CassandraEdgesTable             = flag.String("cassandra.edges_table", "", "Cassandra table name")



)
