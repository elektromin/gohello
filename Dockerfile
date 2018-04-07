FROM alpine:3.1
MAINTAINER Anders Romin <elektromin@hotmail.com>
ADD gohello /usr/bin/gohello
ENTRYPOINT ["gohello"]