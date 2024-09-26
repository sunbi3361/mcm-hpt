#!/bin/bash
cd samples
cd convolution3d
nohup srun -J convolution3d -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/convolution3d.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/convolution3d.err ./convolution3d -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -ni=8192 -nj=8192 -nk=8192 &