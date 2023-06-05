# blosc2

golang (cgo) wrapper for [blosc2](https://github.com/Blosc/c-blosc2)

*a high performance compressor optimized for binary data*

As of today (5th june 2023), no debian libblosc2-dev package is available (there is a libblosc1-dev).
Here in the full install process (can be seen [here](https://github.com/Blosc/c-blosc2) and [here](https://linuxhint.com/install-cmake-on-debian/#2) for cmake)

## Debian c-blosc2 installation

### First install cmake > 3.16

Create and cd to a dev folder then:

```bash
wget https://github.com/Kitware/CMake/releases/download/v3.24.2/cmake-3.24.2.tar.gz
tar -zxvf cmake-3.24.2.tar.gz
cd cmake-3.24.2/
sudo ./bootstrap
sudo make
sudo make install

```

### compile/install c-blosc2 lib and include

Here we install c-blosc2 into /usr/local

```bash
git clone https://github.com/Blosc/c-blosc2
cd c-blosc2
mkdir build
cd build
cmake -DCMAKE_INSTALL_PREFIX=/usr/local ..
cmake --build .
ctest
sudo cmake --build . --target install
```
