#!/bin/bash

configs=("private")

# benchmarks=("convolution2d" "fastwalshtransform" "gups" "jacobi1d" "jacobi2d" "kmeans" "matrixtranspose" "mis" "pagerank" "simpleconvolution" "shoc-reduction" "spmv" "stencil2d" "syrk" "syr2k")
# benchmarks=("gups" "kmeans" "matrixtranspose" "pagerank" "simpleconvolution" "spmv")
# benchmarks=(
#         # 'adi'
#         # 'aes'
#         'atax'
#         'bfs'
#         'bicg'
#         # 'bitonicsort'
#         'blackscholes'
#         'color'
#         'convolution2d'
#         # 'convolution3d'
#         'correlation'
#         'covariance'
#         'doitgen'
#         'fastwalshtransform'
#         'fdtd2d'
#         'fft'
#         'fir'
#         'floydwarshall'
#         'gemm'
#         'gemver'
#         'gesummv'
#         'gramschmidt'
#         'gups'
#         'im2col'
#         # 'jacobi1d'
#         # 'jacobi2d'
#         'kmeans'
#         # 'lenet'
#         'lu'
#         'matrixmultiplication'
#         'matrixtranspose'
#         # 'maxpooling'
#         # 'mineva'
#         'mis'
#         'mm2'
#         'mm3'
#         'mvt'
#         'nbody'
#         'nw'
#         'pagerank'
#         # 'relu'
#         'simpleconvolution'
#         'shoc-reduction'
#         # 'spmv'
#         'sssp'
#         'stencil2d'
#         'syrk'
#         'syr2k'
#         # 'vgg16',
#         # 'xor',
# )

benchmarks=(
        'atax'
        'bicg'
        'color'
        'nbody'
        'gramschmidt'
        'gesummv'
        'floydwarshall'
        'gups'
        'mvt'
        'sssp'
        'syrk'
        'syr2k'
)

for config in ${configs[@]}; 
do
  for benchmark in ${benchmarks[@]}; 
  do
    echo $config $benchmark
    cd $config
    pwd
    bash ${benchmark}.sh
    cd ..
  done
done
        

