<?php
/**
 * @var Goridge\RelayInterface $relay
 */
use Spiral\Goridge;
use Spiral\RoadRunner;

ini_set('display_errors', 'stderr');
require 'vendor/autoload.php';

include 'msg.php';

$rr = new RoadRunner\Worker(new Spiral\Goridge\StreamRelay(STDIN, STDOUT));

while ($body = $rr->receive($context)) {
    try {
    	$msg = new msg();
    	$msg->mergeFromString($body);

        $rr->send($msg->getMsg().$msg->getVersion().(string)$context, (string)$context);
    } catch (\Throwable $e) {
        $rr->error((string)$e);
    }
}
