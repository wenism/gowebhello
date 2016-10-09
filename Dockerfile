FROM scratch
ADD gowebhello /
ADD hello.template.html /
EXPOSE 9999
CMD ["/gowebhello"]