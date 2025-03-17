file_name_root=$1
key_file_name=${file_name_root}.key
crt_file_name=${file_name_root}.crt
key_size=2048
duration=3650
subject="/C=IE/ST=Meath/L=Bettystown/O=Example/OU=Sepa/CN=${file_name_root}"

echo ${key_file_name}


openssl genrsa -out ${key_file_name} ${key_size}
openssl req -new -x509 -sha256 -key ${key_file_name} -out ${crt_file_name} -days ${duration} -subj ${subject}
