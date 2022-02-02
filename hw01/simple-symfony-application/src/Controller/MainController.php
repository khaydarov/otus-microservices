<?php

declare(strict_types=1);

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\Routing\Annotation\Route;

final class MainController extends AbstractController
{
    /**
     * @Route("/", name="main")
     */
    public function index(): JsonResponse
    {
        return $this->json("Hello, dear User");
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
}