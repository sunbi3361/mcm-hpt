#!/bin/bash
cd samples
cd gramschmidt
nohup srun -J gramschmidt -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gramschmidt.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gramschmidt.err ./gramschmidt -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 3000000 -ni=16384 -nj=16384 -k=1 &