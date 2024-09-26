#!/bin/bash
cd samples
cd convolution2d
nohup srun -J convolution2d -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/convolution2d.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/convolution2d.err ./convolution2d -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Ydiv -max-inst 100000000 -ni=32768 -nj=16384 &