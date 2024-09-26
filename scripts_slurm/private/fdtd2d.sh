#!/bin/bash
cd samples
cd fdtd2d
nohup srun -J fdtd2d -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/fdtd2d.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/fdtd2d.err ./fdtd2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -nx=32768 -ny=16384 -max_steps=1 &