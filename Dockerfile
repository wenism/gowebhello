FROM scratch
ADD gowebhello /
ADD hello.template.html /
CMD ["/gowebhello"]