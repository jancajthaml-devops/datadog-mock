# Copyright (c) 2017-2018, Jan Cajthaml <jan.cajthaml@gmail.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM scratch

MAINTAINER Jan Cajthaml <jan.cajthaml@gmail.com>

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="datadog-mock" \
      org.label-schema.description="DataDog mock service Edit" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/jancajthaml/datadog-mock" \
      org.label-schema.vendor="Jan Cajthaml" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

COPY target/datadog_mock /datadog_mock

EXPOSE 8125/UDP

ENTRYPOINT ["/datadog_mock"]
