# xk6_bash_command

you can use it by installing xk6 (k6 extension - https://k6.io/blog/extending-k6-with-xk6/)

example k6 script : 
```aidl
import bash from 'k6/x/bash_command';

export default function () {
        let response = bash.exec("ls",["-la"]);
        console.log(response);
}
```

you can run it with : 
```aidl
./k6 run {script_path}
```