<?php
$url = 'http://dev.test.com/testlgb';
getCurl($url);
function getCurl($url){
    $randIp = mt_rand(0, 255).'.'.mt_rand(0, 255).'.'.mt_rand(0, 255).'.'.mt_rand(0, 255);
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL,$url);
    //curl_setopt($ch, CURLOPT_HTTPHEADER, array('X-FORWARDED-FOR:8.88.88.88', 'CLIENT-IP:123.12.15.18'));  //构造IP

    curl_setopt($ch, CURLOPT_HTTPHEADER, array('X-FORWARDED-FOR:'.$randIp, 'CLIENT-IP:'.$randIp));  //构造IP

    curl_setopt($ch, CURLOPT_REFERER,"http://www.google.com/");   //构造来路
    curl_setopt($ch, CURLOPT_HEADER, 1);
    $out = curl_exec($ch);
    curl_close($ch);

    echo "\n随机IP：".$randIp."\n";
}


function curlPost($url, $post='', $autoFollow=0){
        $ch = curl_init();
        $user_agent = 'Safari Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.73.11 (KHTML, like Gecko) Version/7.0.1 Safari/5
        curl_setopt($ch, CURLOPT_USERAGENT, $user_agent);
        // 2. 设置选项，包括URL
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1); 
        curl_setopt($ch, CURLOPT_HEADER, 0); 
        curl_setopt($ch, CURLOPT_HTTPHEADER, array('X-FORWARDED-FOR:61.135.169.125', 'CLIENT-IP:61.135.169.125'));  //构造IP
        curl_setopt($ch, CURLOPT_REFERER, "http://www.baidu.com/");   //构造来路
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'GET');
        if($autoFollow){
            curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);  //启动跳转链接
            curl_setopt($ch, CURLOPT_AUTOREFERER, true);  //多级自动跳转
        }   
        //  
        if($post!=''){
            curl_setopt($ch, CURLOPT_POST, 1);//post提交方式
            curl_setopt($ch, CURLOPT_POSTFIELDS, $post);
        }   
        // 3. 执行并获取HTML文档内容
        $output = curl_exec($ch);
        curl_close($ch);
        return $output;
    }
