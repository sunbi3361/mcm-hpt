#!/bin/bash
cd samples
cd fft
nohup srun -J fft -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/fft.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/fft.err ./fft -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 1000000000 -MB=1024 -passes=2 &