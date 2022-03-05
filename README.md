# Spacemesh 测试
## gpu-post build
```
cd /home/test/
git clone https://github.com/spacemeshos/gpu-post.git
cd gpu-post
mkdir build
cd build

cmake .. -DSPACEMESHVULKAN=OFF
make


# export CGO_LDFLAGS="-L /home/test/gpu-post/build/src -lgpu-setup"

cp /home/test/gpu-post/build/src/libgpu-setup.so /usr/local/lib/
ldconfig
```

## spacemesh-plotter build
```
cd /home/test
git clone https://github.com/NpoolSpacemesh/spacemesh-plotter.git
cd spacemesh-plotter/
修改plot.go中LabelsPerUnit参数
go build
mkdir plot
把postdata_metadata.json拷贝至plot目录下
time ./spacemesh-plotter
time CUDA_VISIBLE_DEVICES=1 ./spacemesh-plotter
```
