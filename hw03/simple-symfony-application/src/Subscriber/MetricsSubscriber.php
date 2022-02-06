<?php

declare(strict_types=1);

namespace App\Subscriber;

use App\Service\MetricsAdapter;
use Symfony\Component\EventDispatcher\EventSubscriberInterface;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpKernel\Event\ControllerEvent;
use Symfony\Component\HttpKernel\Event\ExceptionEvent;
use Symfony\Component\HttpKernel\Event\ResponseEvent;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;
use Symfony\Component\HttpKernel\KernelEvents;

class MetricsSubscriber implements EventSubscriberInterface
{
    /**
     * @var string
     */
    private const NAMESPACE = 'hw03';

    /**
     * @var array
     */
    private array $metricsRequestCounter = [
        'name' => 'app_request_count',
        'help' => 'Application Request Count',
        'labels' => [
            'method', 'endpoint', 'http_status'
        ]
    ];

    /**
     * @var array
     */
    private array $metricsRequestLatency = [
        'name' => 'app_request_latency_seconds',
        'help' => 'Application Request Latency',
        'labels' => [
            "method", "endpoint"
        ]
    ];

    /**
     * @var MetricsAdapter
     */
    private $metricsAdapter;

    /**
     * @var int
     */
    private $startTime;

    public function __construct(MetricsAdapter $metricsAdapter)
    {
        $this->metricsAdapter = $metricsAdapter;
    }

    /**
     * @param ControllerEvent $event
     */
    public function onKernelController(): void
    {
        $this->startTime = time();
    }

    /**
     * @throws \Prometheus\Exception\MetricsRegistrationException
     */
    public function onKernelException(ExceptionEvent $event)
    {
        $throwable = $event->getThrowable();
        if ($throwable instanceof NotFoundHttpException) {
            $response = new JsonResponse([
                'code' => 404,
                'message' => 'Route not found'
            ]);
        } else {
            $response = new JsonResponse([
                'code' => 500,
                'message' => $throwable->getMessage()
            ]);
        }

        $event->setResponse($response);
    }

    /**
     * @param ResponseEvent $event
     * @throws \Prometheus\Exception\MetricsRegistrationException
     */
    public function onKernelResponse(ResponseEvent $event): void
    {
        $request = $event->getRequest();
        $response = $event->getResponse();
        $endTime = time() - $this->startTime;

        /** Calc request count */
        $requestCounter = $this->metricsAdapter->getPrometheusRegistry()->getOrRegisterCounter(
            self::NAMESPACE,
            $this->metricsRequestCounter['name'],
            $this->metricsRequestCounter['help'],
            $this->metricsRequestCounter['labels']
        );

        $requestCounter->inc([
            $request->getMethod(),
            $request->getRequestUri(),
            $response->getStatusCode()
        ]);

        /** Calc request latency */
        $requestLatencyHistogram = $this->metricsAdapter->getPrometheusRegistry()->getOrRegisterHistogram(
            self::NAMESPACE,
            $this->metricsRequestLatency['name'],
            $this->metricsRequestLatency['help'],
            $this->metricsRequestLatency['labels']
        );

        $requestLatencyHistogram->observe($endTime, [
            $request->getMethod(),
            $request->getRequestUri()
        ]);
    }

    /**
     * @inheritdoc
     */
    public static function getSubscribedEvents(): array
    {
        return [
            KernelEvents::CONTROLLER => 'onKernelController',
            KernelEvents::EXCEPTION => 'onKernelException',
            KernelEvents::RESPONSE => 'onKernelResponse'
        ];
    }
}