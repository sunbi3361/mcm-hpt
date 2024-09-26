#!/usr/bin/python3

# configs = ['private', 'shared', 'mgvm', 'mgvm-nobalance']
configs = ['private']

# benchmarks = [
#         'convolution2d',
#         'fastwalshtransform',
#         'gups',
#         'jacobi1d',
#         'jacobi2d',
#         'kmeans',
#         'matrixtranspose',
#         'mis',
#         'pagerank',
#         'simpleconvolution',
#         'shoc-reduction',
#         'spmv',
#         'stencil2d',
#         'syrk',
#         'syr2k',
#         ]

# benchmarks = [
#         # 'convolution2d',
#         'fastwalshtransform',
#         'gups',
#         # 'jacobi1d',
#         # 'jacobi2d',
#         'kmeans',
#         'matrixtranspose',
#         'mis',
#         'pagerank',
#         'simpleconvolution',
#         'shoc-reduction',
#         'spmv',
#         'stencil2d',
#         'syrk',
#         'syr2k',
#         'bfs',
#         ]

benchmarks = [
        # 'adi',  
        # 'aes', 
        'atax', 
        'bfs', #
        'bicg', 
        # 'bitonicsort', 
        'blackscholes',
        'color',
        'convolution2d',
        # 'convolution3d', #
        'correlation',
        'covariance',
        'doitgen',
        'fastwalshtransform', #
        'fdtd2d',
        'fft',
        'fir',
        'floydwarshall',
        'gemm',
        'gemver',
        'gesummv',
        'gramschmidt',
        'gups', #
        'im2col',
        # 'jacobi1d', #
        # 'jacobi2d', #
        'kmeans', #
        # 'lenet',
        'lu',
        'matrixmultiplication',
        'matrixtranspose', #
        # 'maxpooling',
        # 'mineva',
        'mis', #
        'mm2',
        'mm3',
        'mvt',
        'nbody',
        'nw',
        'pagerank', #
        # 'relu',
        'simpleconvolution', #
        'shoc-reduction', #
        'spmv', #
        'sssp',
        'stencil2d', #
        'syrk', #
        'syr2k', #
        # 'vgg16',
        # 'xor',
        ]

for config in configs:
    for benchmark in benchmarks:
        print(config, benchmark)
        submit_file_name = config + '/' + benchmark + ".sh"
        submit_file = open(submit_file_name, "w")
        submit_file.write("#!/bin/bash\n")
        submit_file.write("cd samples\n")
        submit_file.write("cd " + benchmark + "\n")
        submit_file.write("nohup srun -J " + benchmark + " -w compasslab5" + " ")
        submit_file.write("--output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/" + config + "/output/" + benchmark + ".out" + " ")
        submit_file.write("--error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/" + config + "/output/" + benchmark + ".err" + " ")
        submit_file.write("./" + benchmark + " ")
        submit_file.write("-timing ")
        submit_file.write("-no-progress-bar ")
        submit_file.write("-report-all ")
        submit_file.write("-scheduling lasp ")

        if config == 'private':
            submit_file.write("-platform-type privatetlb ")
            submit_file.write("-mem-allocator-type lasp ")
            submit_file.write("-use-lasp-mem-alloc ")
        elif config == 'shared':
            submit_file.write("-platform-type xortlb ")
            submit_file.write("-mem-allocator-type lasp ")
            submit_file.write("-use-lasp-mem-alloc ")
            submit_file.write("-l2-tlb-striping 1 ")
        elif config == 'mgvm':
            submit_file.write("-platform-type customtlb ")
            submit_file.write("-use-lasp-hsl-mem-alloc ")
            submit_file.write("-switch-tlb-striping ")
        elif config == 'mgvm-nobalance':
            submit_file.write("-platform-type customtlb ")
            submit_file.write("-use-lasp-hsl-mem-alloc ")

        if benchmark == 'convolution2d':
            submit_file.write("-sched-partition Ydiv ")
        elif benchmark == 'fastwalshtransform':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'gups':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'jacobi1d':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'jacobi2d':
            submit_file.write("-sched-partition Ydiv ")
        elif benchmark == 'kmeans':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'matrixtranspose':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'mis':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'pagerank':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'simpleconvolution':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'shoc-reduction':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'spmv':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'stencil2d':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'syrk':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'syr2k':
            submit_file.write("-sched-partition Xdiv ")
        elif benchmark == 'bfs':
            submit_file.write("-sched-partition Xdiv ")
        else:
            submit_file.write("-sched-partition Xdiv ")

        # set appropriate HSL values
        if config == 'mgvm' or config == 'mgvm-nobalance':
            if benchmark == 'convolution2d':
                submit_file.write("-custom-hsl 16384 ")
                submit_file.write("-mem-allocator-type hslaware-32 ")
            if benchmark == 'fastwalshtransform':
                submit_file.write("-custom-hsl 2048 ")
                submit_file.write("-mem-allocator-type hslaware-4 ")
            if benchmark == 'gups':
                submit_file.write("-custom-hsl 1024 ")
                submit_file.write("-mem-allocator-type hslaware-2 ")
            if benchmark == 'jacobi1d':
                submit_file.write("-custom-hsl 16384 ")
                submit_file.write("-mem-allocator-type hslaware-32 ")
            if benchmark == 'jacobi2d':
                submit_file.write("-custom-hsl 4096 ")
                submit_file.write("-mem-allocator-type hslaware-8 ")
            if benchmark == 'kmeans':
                submit_file.write("-custom-hsl 4096 ")
                submit_file.write("-mem-allocator-type hslaware-8 ")
            if benchmark == 'matrixtranspose':
                submit_file.write("-custom-hsl 1024 ")
                submit_file.write("-mem-allocator-type hslaware-2 ")
            if benchmark == 'mis':
                submit_file.write("-custom-hsl 512 ")
                submit_file.write("-mem-allocator-type hslaware-1 ")
            if benchmark == 'pagerank':
                submit_file.write("-custom-hsl 8192 ")
                submit_file.write("-mem-allocator-type hslaware-16 ")
            if benchmark == 'simpleconvolution':
                submit_file.write("-custom-hsl 16384 ")
                submit_file.write("-mem-allocator-type hslaware-32 ")
            if benchmark == 'shoc-reduction':
                submit_file.write("-custom-hsl 16384 ")
                submit_file.write("-mem-allocator-type hslaware-32 ")
            if benchmark == 'spmv':
                submit_file.write("-custom-hsl 512 ")
                submit_file.write("-mem-allocator-type hslaware-1 ")
            if benchmark == 'stencil2d':
                submit_file.write("-custom-hsl 1024 ")
                submit_file.write("-mem-allocator-type hslaware-2 ")
            if benchmark == 'syrk':
                submit_file.write("-custom-hsl 1024 ")
                submit_file.write("-mem-allocator-type hslaware-2 ")
            if benchmark == 'syr2k':
                submit_file.write("-custom-hsl 512 ")
                submit_file.write("-mem-allocator-type hslaware-1 ")

        # limit super long benchmarks
        if benchmark == 'adi': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'aes': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'atax': 
            submit_file.write("-max-inst 3000000 ")
        if benchmark == 'bfs': 
            submit_file.write("-max-inst 300000000 ")
        if benchmark == 'bicg': 
            submit_file.write("-max-inst 3000000 ")
        if benchmark == 'bitonicsort': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'color': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'convolution2d':
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'convolution3d':
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'correlation': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'covariance': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'doitgen': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'fastwalshtransform': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'fdtd2d': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'fft': 
            submit_file.write("-max-inst 1000000000 ")
        if benchmark == 'fir': 
            submit_file.write("-max-inst 1000000000 ")
        if benchmark == 'floydwarshall': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'gemm': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'gemver': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'gesummv': 
            submit_file.write("-max-inst 3000000 ")
        if benchmark == 'gramschmidt': 
            submit_file.write("-max-inst 3000000 ")
        if benchmark == 'gups': 
            submit_file.write("-max-inst 3000000 ")
        if benchmark == 'im2col': 
            submit_file.write("-max-inst 1000000000 ")
        if benchmark == 'jacobi1d':
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'jacobi2d':
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'kmeans': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'lenet':  
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'lu': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'matrixmultiplication': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'matrixtranspose': 
            submit_file.write("-max-inst 10000000 ")
        if benchmark == 'maxpooling': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'mineva': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'mis': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'mm2': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'mm3': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'mvt': 
            submit_file.write("-max-inst 3000000 ")
        if benchmark == 'nbody': 
            submit_file.write("-max-inst 1000000000 ")
        if benchmark == 'nw': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'pagerank': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'relu': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'simpleconvolution': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'shoc-reduction': 
            submit_file.write("-max-inst 1000000000 ")
        if benchmark == 'spmv': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'sssp': 
            submit_file.write("-max-inst 300000000 ")
        if benchmark == 'stencil2d': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'syrk': 
            submit_file.write("-max-inst 3000000 ")
        if benchmark == 'syr2k': 
            submit_file.write("-max-inst 3000000 ")
        if benchmark == 'vgg16': 
            submit_file.write("-max-inst 100000000 ")
        if benchmark == 'xor': 
            submit_file.write("-max-inst 100000000 ")

        # set benchmark specific parameters
        if benchmark == 'adi': # 
            submit_file.write("-n=131072 -steps=1 ")
        if benchmark == 'aes': # 
            submit_file.write("-length=65536 ")
        if benchmark == 'atax': # 4096
            submit_file.write("-x=32768 -y=32768 ")
        if benchmark == 'bfs': # 4096
            submit_file.write("-node 32768 -degree 32768 -depth 0 ")
        if benchmark == 'bicg': # 4096
            submit_file.write("-x=32768 -y=32768 ")
        if benchmark == 'bitonicsort': #
            submit_file.write("-length=131072 -order-asc=true ")
        if benchmark == 'blackscholes': # 3072
            submit_file.write("-width=8192 -height=8192 ")
        if benchmark == 'color': # 3584
            submit_file.write("-numNodes=134217728 -numItems=268435456 ")
        if benchmark == 'convolution2d': # 4096
            submit_file.write("-ni=32768 -nj=16384 ")
        if benchmark == 'convolution3d': # 4096
            submit_file.write("-ni=8192 -nj=8192 -nk=8192 ")
        if benchmark == 'correlation': # 4096
            submit_file.write("-num_m=16384 -num_n=32768 ")
        if benchmark == 'covariance': # 5376
            submit_file.write("-num_m=24576 -num_n=32768 ")
        if benchmark == 'doitgen': # 4096
            submit_file.write("-r=1024 -q=1024 -p=512 ")
        if benchmark == 'fastwalshtransform': # 3686
            submit_file.write("-length=966367641 ")
        if benchmark == 'fdtd2d': # 6144
            submit_file.write("-nx=32768 -ny=16384 -max_steps=1 ")
        if benchmark == 'fft': # 1024
            submit_file.write("-MB=1024 -passes=2 ")
        if benchmark == 'fir': # 
            submit_file.write("-length=536870912 ")
        if benchmark == 'floydwarshall': # 2048
            submit_file.write("-node=16384 -iter=0 ")
        if benchmark == 'gemm': # 3072
            submit_file.write("-num_i=16384 -num_j=16384 -num_k=16384 ")
        if benchmark == 'gemver': # 4097
            submit_file.write("-n=32768 ")
        if benchmark == 'gesummv': # 4608
            submit_file.write("-n=24576 ")
        if benchmark == 'gramschmidt': # 3072
            submit_file.write("-ni=16384 -nj=16384 -k=1 ")
        if benchmark == 'gups': # 4096
            submit_file.write(" ")
        if benchmark == 'im2col': # 3072
            submit_file.write("-N=16 -C=3 -H=1024 -W=1024 -kernel-height=3 -kernel-width=3 ")
        if benchmark == 'jacobi1d': # 512*2
            submit_file.write("-n=134217728 -steps=1 ")
        if benchmark == 'jacobi2d': # 128*4
            submit_file.write("-n=32768 -steps=1 ")
        if benchmark == 'kmeans': # 128*8
            submit_file.write("-points=4194304 -features=128 -clusters=20 -max-iter=1 ")
        if benchmark == 'lenet': # 
            submit_file.write("-epoch=1 -batch-size=32 ")
        if benchmark == 'lu': # 4096
            submit_file.write("-n=32768 -k=1 ")
        if benchmark == 'matrixmultiplication': # 5120
            submit_file.write("-x=32768 -y=16384 -z=16384 ")
        if benchmark == 'matrixtranspose': # 4608
            submit_file.write("-width=24576 ")
        if benchmark == 'maxpooling': # 
            submit_file.write("-n=1 -c=1 -h=32 -w=32 -kernel-h=3 -kernel-w=3 -pad-h=0 -pad-w=0 -stride-h=1 -stride-w=1")
        if benchmark == 'mineva': # 
            submit_file.write("-epoch=1 -max-batch-per-epoch=2 -batch-size=1024")
        if benchmark == 'mis': # 3080
            submit_file.write("-numNodes=134217728 -numItems=2097152 ")
        if benchmark == 'mm2': # 2048
            submit_file.write("-ni=16384 -nj=8192 -nk=8192 -nl=8192 ")
        if benchmark == 'mm3': # 3840
            submit_file.write("-ni=16384 -nj=16384 -nk=8192 -nl=8192 -nm=8192 ")
        if benchmark == 'mvt': # 4096
            submit_file.write("-n=32768 ")
        if benchmark == 'nbody': # 2048
            submit_file.write("-iter=8 -particles=33554432 ")
        if benchmark == 'nw': # 2304
            submit_file.write("-length=24576 ")
        if benchmark == 'pagerank': # 256*4
            submit_file.write("-node=16384 -sparsity=0.5 -iterations=1 ")
        if benchmark == 'relu': # 
            submit_file.write("-length=262144 ")
        if benchmark == 'simpleconvolution': # 4096
            submit_file.write("-width=16384 -height=32760 ")
        if benchmark == 'shoc-reduction': # 4096
            submit_file.write("-Size=536870912 -Iterations=2 ")
        if benchmark == 'spmv': # 
            submit_file.write("-dim=8388608 -sparsity=0.00001 ")
        if benchmark == 'sssp': # 3584
            submit_file.write("-numNodes=134217728 -numItems=268435456 ")
        if benchmark == 'stencil2d': # 4100
            submit_file.write("-row=32768 -col=16384 ")
        if benchmark == 'syrk': # 6144
            submit_file.write("-ni=32768 -nj=16384 ")
        if benchmark == 'syr2k': # 3072
            submit_file.write("-ni=16384 -nj=16384 ")
        if benchmark == 'vgg16': # 
            submit_file.write("-epoch=1 -max-batch-per-epoch=2 -batch-size=8 ")
        if benchmark == 'xor': # 
            submit_file.write(" ")
    
        submit_file.write("&")

