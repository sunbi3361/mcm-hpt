#!/bin/bash
cd samples
cd covariance
nohup srun -J covariance -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/covariance.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/covariance.err ./covariance -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -num_m=24576 -num_n=32768 &