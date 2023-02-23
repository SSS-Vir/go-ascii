
# Build
```bash
git clone https://github.com/SSS-Vir/go-ascii.git
cd go-ascii
go get go-ascii
go build .
```
#  Usage
```bash
.\go-ascii.exe --help
```
Usage:<br>
&emsp;-file string<br>
        path to file in quotes e.g. ".\path\to\file"<br>
&emsp;-fps uint<br>
&emsp;&emsp;non negative num (default 18)<br>
&emsp;-resample string<br>
&emsp;One of [box catmullrom bartlett lanczos bspline gaussian hann nearestneighbor linear hermite mitchellnetravali hamming blackman welch cosine] (default "nearestneighbor")<br>
&emsp;-size string<br>
&emsp;&emsp;non negative WIDTHxHEIGHT
