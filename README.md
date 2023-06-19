Stock exchange simulator to match offers for purchases and sales of shares where:
- Need to be minimally performative
- You will need to work with "in memory" operations (because the simulation operations need to be in real time, and so that it is not necessary to query and lock the database at all times)
- Major allocations will need to stay on the heap and not on the stack