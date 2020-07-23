ls -l | awk '
BEGIN { n=2 }
n == 1 { print $0; n=2; next }
n == 2 { n=1; next }'
