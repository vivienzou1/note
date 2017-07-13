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
