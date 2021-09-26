# Golang results

Testing command: `wrk -c 500 -d 30s -t 2`

| Framework | Server | Requests per second | Requests per core per second | Requests per dollar per second |
|-----------|--------|---------------------|------------------------------|--------------------------------|
| Echo | DO Premium Intel 2GB 1vCPU ($12/m) | 22755.99 | 22755.99 | 1896.33 |
| Fiber | DO Premium Intel 2GB 1vCPU ($12/m) | 22444.36 | 22444.36 | 1870.36 |
| Gin | DO Premium Intel 2GB 1vCPU ($12/m) | 20459.36 | 20459.36 | 1704.94 |
| net/http | DO Premium Intel 2GB 1vCPU ($12/m) | 19248.31 | 19248.31 | 1604.02 |
| Echo | DO Premium Intel 2GB 2vCPU ($18/m) | 39854.03 | 19927.015 | 2214.11 |
| Fiber | DO Premium Intel 2GB 2vCPU ($18/m) | 37992.13 | 18996.065 | 2110.67 |
| Gin | DO Premium Intel 2GB 2vCPU ($18/m) | 20220.60 | 10110.3 | 1123.36 |
| net/http | DO Premium Intel 2GB 2vCPU ($18/m) | 17791.29 | 8895.645 | 988.40 |
| Echo | Vultr High-Frequency 2GB 1vCPU ($12/m) | 30612.62 | 30612.62 | 2551.05 |
| Fiber | Vultr High-Frequency 2GB 1vCPU ($12/m) | 58270.56 | 58270.56 | 4855.88 |
| Gin | Vultr High-Frequency 2GB 1vCPU ($12/m) | 29844.04 | 29844.04 | 2487 |
| net/http | Vultr High-Frequency 2GB 1vCPU ($12/m) | 31813.91 | 31813.91 | 2651.15 |
| Echo | Vultr High-Frequency 2GB 2vCPU ($18/m) | 61919.90 | 30959.95 | 3439.99 |
| Fiber | Vultr High-Frequency 2GB 2vCPU ($18/m) | 140893.29 | 70446.645 | 7827.40 |
| Gin | Vultr High-Frequency 2GB 2vCPU ($18/m) | 60032.24 | 30016.12 | 3335.12 |
| net/http | Vultr High-Frequency 2GB 2vCPU ($18/m) | 64038.93 | 32019.465 | 3557.71 |
| Echo | Vultr Bare Metal E3-1270v6 ($120/m)* | 230098.39 | 28762.29 | 1917.48 |
| Fiber | Vultr Bare Metal E3-1270v6 ($120/m)* | 268324.87 | 33540.608 | 2236.04 | 
| Gin | Vultr Bare Metal E3-1270v6 ($120/m)* | 228151.05 | 28518.88 | 1901.25 |
| net/http | Vultr Bare Metal E3-1270v6 ($120/m)* | 234306.43 | 29288.30 | 1952.55 |

asterisk - wrk was running on the same machine