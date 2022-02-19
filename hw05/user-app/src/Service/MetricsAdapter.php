<?php

declare(strict_types=1);

namespace App\Service;

use Prometheus\CollectorRegistry;
use Prometheus\Storage\APC;

class MetricsAdapter
{
    public function __construct()
    {
        $this->prometheusRegistry = new CollectorRegistry(new APC());
    }

    /**
     * @return CollectorRegistry
     */
    public function getPrometheusRegistry(): CollectorRegistry
    {
        return $this->prometheusRegistry;
    }
}