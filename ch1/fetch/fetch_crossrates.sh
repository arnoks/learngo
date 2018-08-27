#!/bin/bash
mkdir -p ~/xrates

./fetch https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml?b35919cb820e785ee62462bc73a5f4d1  > ~/xrates/crossrates.xml
