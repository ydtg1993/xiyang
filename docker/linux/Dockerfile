FROM centos:7
ENV container docker
RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == \
systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;

RUN curl -L -o go.tar.gz https://golang.org/dl/go1.17.6.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm -rf go.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

VOLUME [ "/sys/fs/cgroup" ]
CMD ["/usr/sbin/init"]
COPY /chrome.rpm  /home/chrome.rpm

RUN yum -y install /home/chrome.rpm


# set timezome
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
