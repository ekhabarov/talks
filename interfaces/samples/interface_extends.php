<?
interface a {
    public function foo();
}

interface b {
    public function bar();
}


class ab implements a, b
{
    public function foo() { }
    public function bar() { }
}
?>
