gen --sqltype=mysql \
--connstr "cx_rw:cx_rw@2020@@tcp(172.17.19.131:3306)/yx" \
--database yx  \
--json \
--gorm \
--out ./example \
--module github.com/xiashauang/demitasse.cn \
--overwrite