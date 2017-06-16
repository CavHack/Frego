# Frego

QUEGEL invarant of Pregel for BiBFS on Distributed Cluster Graphs. 

"Pioneered by Google's Pregel, many distributed systems have been developed for large-scale graph analytics. 
These systems expose the user-friendly "think like a vertex" programming interface to users, and exhibit good horizontal scalability.
However, these systems are designed for tasks where the majority of graph vertices participate in computation, but are not suitable for processing
light-workload graph queries where only a small fraction of vertices need to be accessed. 

The programming paradigm adopted by these systems can seriously under-utilize the resources in a cluster for graph query processing. In this work, we develop a new open-source system, called Quegel, 
for querying big graphs, which treats queries as first-class citizens in the design of its computing model. 

Users only need to specify the Pregel-like algorithm for a generic query, and Quegel processes light-workload graph queries on demand using a novel superstep-sharing 
execution model to effectively utilize the cluster resources. Quegel further provides 
a convenient interface for constructing graph indexes, which significantly improve query performance 
but are not supported by existing graph-parallel systems. 

Our experiments verified that Quegel is highly efficient in answering various types of graph queries and is up 
to orders of magnitude faster than existing systems."

https://arxiv.org/abs/1601.06497
