# # NOTE: registry keys for IE 8, may vary for other versions
# $regPath = 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings'

# function Clear-Proxy
# {
#     Set-ItemProperty -Path $regPath -Name ProxyEnable -Value 0
#     Set-ItemProperty -Path $regPath -Name ProxyServer -Value ''
#     Set-ItemProperty -Path $regPath -Name ProxyOverride -Value ''

#     [Environment]::SetEnvironmentVariable('http_proxy', $null, 'User')
#     [Environment]::SetEnvironmentVariable('https_proxy', $null, 'User')
# }

# function Set-Proxy
# {
#     $proxy = 'http://example.com'

#     Set-ItemProperty -Path $regPath -Name ProxyEnable -Value 1
#     Set-ItemProperty -Path $regPath -Name ProxyServer -Value $proxy
#     Set-ItemProperty -Path $regPath -Name ProxyOverride -Value '<local>'

#     [Environment]::SetEnvironmentVariable('http_proxy', $proxy, 'User')
#     [Environment]::SetEnvironmentVariable('https_proxy', $proxy, 'User')
# }