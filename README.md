# nautilus - 鹦鹉螺

### 镜像打包到aliyun

~~~
$ docker build --rm -t nautilus .

$ docker login --username=glauca registry.cn-hangzhou.aliyuncs.com

$ docker pull registry.cn-hangzhou.aliyuncs.com/glauca/nautilus:[镜像版本号]

$ docker login --username=glauca registry.cn-hangzhou.aliyuncs.com
$ docker tag [ImageId] registry.cn-hangzhou.aliyuncs.com/glauca/nautilus:[镜像版本号]
$ docker push registry.cn-hangzhou.aliyuncs.com/glauca/nautilus:[镜像版本号]
~~~
