<?php

declare(strict_types=1);

namespace App\Service;

use Prometheus\CollectorRegistry;
use Prometheus\Storage\Redis;

class MetricsAdapter
{
    public function __construct()
    {
        $options = [
            'host' => $_ENV['REDIS_HOST'],
            'port' => 6379,
            'timeout' => 0.1, // in seconds
            'read_timeout' => '10', // in seconds
            'persistent_connections' => false
        ];

        if (!empty($_ENV['REDIS_PASS'])) {
            $options['password'] = $_ENV['REDIS_PASS'];
        }

        \Prometheus\Storage\Redis::setDefaultOptions($options);
        $this->prometheusRegistry = new CollectorRegistry(new Redis());
    }

    /**
     * @return CollectorRegistry
     */
    public function getPrometheusRegistry(): CollectorRegistry
    {
        return $this->prometheusRegistry;
    }
}