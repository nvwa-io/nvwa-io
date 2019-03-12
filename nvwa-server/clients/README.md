

## Install jenkins

### Centos

```
sudo yum install wget -y
sudo yum install java -y

# get jenkins repo
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo

# install jenkins public key
sudo rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io.key

# install jenkins
sudo yum install jenkins

```

Check jenkins's install path
```
rpm -ql jenkins
```

Jenkins's workspace paths

```
/usr/lib/jenkins/：jenkins安装目录，war包会放在这里。
/etc/sysconfig/jenkins：jenkins配置文件，“端口”，“JENKINS_HOME”等都可以在这里配置。
/var/lib/jenkins/：默认的JENKINS_HOME。
/var/log/jenkins/jenkins.log：jenkins日志文件。
```


Boot jenkins

```
# start jenkins
sudo systemctl start jenkins

# check jenkins status
sudo systemctl status jenins
```

View jenkins in browser `http://localhost:8080`

Get `Administrator password` in file: `/var/lib/jenkins/secrets/initialAdminPassword`

```
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
```

安装推荐的 Jenkins 插件。