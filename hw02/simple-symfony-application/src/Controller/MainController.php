<?php

declare(strict_types=1);

namespace App\Controller;

use Prometheus\CollectorRegistry;
use Prometheus\RenderTextFormat;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\HttpFoundation\JsonResponse;

class MainController extends AbstractController
{
    /**
     * @var CollectorRegistry
     */
    private CollectorRegistry $prometheusRegistry;

    public function __construct()
    {
        $this->prometheusRegistry = CollectorRegistry::getDefault();
    }

    /**
     * @Route("/", name="index")
     */
    public function index(): JsonResponse
    {
        return $this->json([
            'Hello!'
        ]);
    }

    /**
     * @Route("/health", name="health")
     */
    public function health(): JsonResponse
    {
        return $this->json([
            'status' => 'OK!'
        ]);
    }

    /**
     * @Route("/metrics", name="metrics")
     */
    public function metrics(): Response
    {
        $renderer = new RenderTextFormat();
        $result = $renderer->render($this->prometheusRegistry->getMetricFamilySamples());

        return new Response($result);
    }
}