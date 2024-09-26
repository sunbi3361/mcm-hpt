#!/bin/bash
cd samples
cd doitgen
nohup srun -J doitgen -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/doitgen.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/doitgen.err ./doitgen -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -r=1024 -q=1024 -p=512 &