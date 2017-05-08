 public function check() {
        $getfilter    = "'|(and|or)\\b.+?(>|<|=|in|like)|\\/\\*.+?\\*\\/|<\\s*script\\b|\\bEXEC\\b|UNION.+?SELECT|UPDATE.+?SET|INSERT\\s+INTO.+?VALUES|(SELECT|DELETE).+?FROM|(CREATE|ALTER|DROP|TRUNCATE)\\s+(TABLE|DATABASE)";
        $postfilter   = "\\b(and|or)\\b.{1,6}?(=|>|<|\\bin\\b|\\blike\\b)|\\/\\*.+?\\*\\/|<\\s*script\\b|\\bEXEC\\b|UNION.+?SELECT|UPDATE.+?SET|INSERT\\s+INTO.+?VALUES|(SELECT|DELETE).+?FROM|(CREATE|ALTER|DROP|TRUNCATE)\\s+(TABLE|DATABASE)";
        $cookiefilter = "\\b(and|or)\\b.{1,6}?(=|>|<|\\bin\\b|\\blike\\b)|\\/\\*.+?\\*\\/|<\\s*script\\b|\\bEXEC\\b|UNION.+?SELECT|UPDATE.+?SET|INSERT\\s+INTO.+?VALUES|(SELECT|DELETE).+?FROM|(CREATE|ALTER|DROP|TRUNCATE)\\s+(TABLE|DATABASE)";

        foreach ( $_GET as $key => $value ) {
            $this->StopAttack( $key, $value, $getfilter );
        }
        foreach ( $_POST as $key => $value ) {
            $this->StopAttack( $key, $value, $postfilter );
        }
        foreach ( $_COOKIE as $key => $value ) {
            $this->StopAttack( $key, $value, $cookiefilter );
        }
    }

    private function StopAttack( $StrFiltKey, $StrFiltValue, $ArrFiltReq ) {

        if ( is_array( $StrFiltValue ) ) {
            $StrFiltValue = implode( $StrFiltValue );
        }
        if ( preg_match( "/" . $ArrFiltReq . "/is", $StrFiltValue ) == 1 ) {
            $msg = "<br><br>操作IP: " . getIp() . "<br>操作时间: " . strftime( "%Y-%m-%d %H:%M:%S" ) . "<br>操作页面:" . $_SERVER["PHP_SELF"] . "<br>提交方式: " . $_SERVER["REQUEST_METHOD"] . "<br>提交参数: " . $StrFiltKey . "<br>提交数据: " . $StrFiltValue );
            echo  "非法的网站请求!";
        }
    }
    
     //记录并发邮件
    private function slog( $logs ) {
    
    }
