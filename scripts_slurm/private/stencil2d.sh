#!/bin/bash
cd samples
cd stencil2d
nohup srun -J stencil2d -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/stencil2d.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/stencil2d.err ./stencil2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -row=32768 -col=16384 &