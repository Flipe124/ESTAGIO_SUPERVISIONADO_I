<?php

require_once './IFinance.php';
require_once './Finance.php';

class InputFinance extends Finance implements IFinance {

    public $registrationDate;
    public function get() {
        // SELECT return.
    }

}
