<?php

class Task {

	# ********** Variable Declaration **********

	private $description;
	private $value;

	# ********** Variable Functions **********

	public function getDescription() {
		return $this->description;
	}

	public function setDescription($description) {
		if (! is_string($description))
			return throw new InvalidArgumentException('Função só aceita argumentos do tipo string!');
		$this->description = $description;
	}

	public function getValue() {
		return $this->value;
	}

	public function setValue($value) {
		if (! is_float((float) $value))
			return throw new InvalidArgumentException('Função só aceita argumentos dos tipos inteiro/real!');
		$this->value = $value;
	}

}

# >>> Start tests !

$task = new Task();
$task->setDescription('hello world');
echo $task->getDescription();

?>
