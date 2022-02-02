<?php

declare(strict_types=1);

namespace App\Controller;

use App\Entity\User;
use App\Repository\UserRepository;
use Prometheus\CollectorRegistry;
use Prometheus\Counter;
use Prometheus\RenderTextFormat;
use Prometheus\Storage\InMemory;
use Redis;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

/**
 * Class UserController
 *
 * @package App\Controller
 */
final class UserController extends AbstractController
{
    private $metricsMap = [
        'namespace' => 'otus_hw02',
        'getUserAction' => [
            'rps' => [
                'name' => 'get_user_rps',
                'help' => 'get User endpoint RPS',
                'labels' => []
            ],
            'errors' => [
                'name' => 'get_user_errors',
                'help' => 'get User endpoint errors',
                'labels' => []
            ]
        ],
        'postUserAction' => [
            'rps' => [
                'name' => 'post_user_rps',
                'help' => 'New User creation endpoint RPS',
                'labels' => []
            ],
            'errors' => [
                'name' => 'post_user_errors',
                'help' => 'New User creation errors',
                'labels' => [
                    'code'
                ]
            ]
        ],
        'putUserAction' => [
            'rps' => [
                'name' => 'put_user_rps',
                'help' => 'User update endpoint RPS',
                'labels' => []
            ],
            'errors' => [
                'name' => 'put_user_errors',
                'help' => 'User update endpoint errors',
                'labels' => []
            ]
        ],
        'deleteUserAction' => [
            'rps' => [
                'name' => 'delete_user_rps',
                'help' => 'Delete User endpoint RPS',
                'labels' => []
            ],
            'errors' => [
                'name' => 'delete_user_errors',
                'help' => 'Delete User endpoint errors',
                'labels' => []
            ]
        ]
    ];

    /**
     * @var UserRepository
     */
    private UserRepository $userRepository;

    /**
     * @var CollectorRegistry
     */
    private CollectorRegistry $prometheusRegistry;

    public function __construct(UserRepository $userRepository)
    {
        $this->userRepository = $userRepository;
        $this->prometheusRegistry = CollectorRegistry::getDefault();
    }

    /**
     * @Route("/user", name="postUserAction", methods={"POST"})
     *
     * @param Request $request
     *
     * @return JsonResponse
     */
    public function postUserAction(Request $request): JsonResponse
    {
        $rpsCounter = $this->getApiMethodMetricsCounter(__FUNCTION__, 'rps');
        $rpsCounter->inc();

        $requestData = json_decode($request->getContent(), true);

        $username = $requestData['username'] ?? '';
        $firstName = $requestData['firstName'] ?? '';
        $lastName = $requestData['lastName'] ?? '';
        $email = $requestData['email'] ?? '';
        $phone = $requestData['phone'] ?? '';

        try {
            $user = new User(
                $this->userRepository->nextIdentity(),
                $username,
                $firstName,
                $lastName,
                $email,
                $phone
            );

            $this->userRepository->insert($user);

            return $this->json([
                'id' => $user->getId()
            ]);
        } catch (\Throwable $e) {
            $errorsCounter = $this->getApiMethodMetricsCounter(__FUNCTION__, 'errors');
            $errorsCounter->inc();

            return $this->json([
                'code' => 0,
                'message' => $e->getMessage()
            ]);
        }
    }

    /**
     * @Route("/user/{id<\d+>}", name="getUserAction", methods={"GET"})
     *
     * @param int $id
     *
     * @return JsonResponse
     */
    public function getUserAction(int $id): JsonResponse
    {
        $rpsCounter = $this->getApiMethodMetricsCounter(__FUNCTION__, 'rps');
        $rpsCounter->inc();

        try {
            $user = $this->userRepository->findById($id);

            if ($user === null) {
                return $this->json([
                    'code' => 0,
                    'message' => 'User not found'
                ]);
            }

            return $this->json([
                'id' => $id,
                'username' => $user->getUsername(),
                'firstName' => $user->getFirstName(),
                'lastName' => $user->getLastName(),
                'email' => $user->getEmail(),
                'phone' => $user->getPhone()
            ]);
        } catch (\Throwable $e) {
            $errorsCounter = $this->getApiMethodMetricsCounter(__FUNCTION__, 'errors');
            $errorsCounter->inc();

            return $this->json([
                'code' => 500,
                'message' => $e->getMessage()
            ]);
        }
    }

    /**
     * @Route("/user/{id<\d+>}", name="putUserAction", methods={"PUT"})
     *
     * @param Request $request
     *
     * @return JsonResponse
     */
    public function putUserAction(int $id, Request $request): JsonResponse
    {
        $rpsCounter = $this->getApiMethodMetricsCounter(__FUNCTION__, 'rpc');
        $rpsCounter->inc();

        try {
            $user = $this->userRepository->findById($id);
            if ($user === null) {
                return $this->json([
                    'code' => 404,
                    'message' => 'User not found'
                ]);
            }

            $requestData = json_decode($request->getContent(), true);
            $firstName = $requestData['firstName'] ?? '';
            $lastName = $requestData['lastName'] ?? '';
            $email = $requestData['email'] ?? '';
            $phone = $requestData['phone'] ?? '';

            if (!empty($firstName)) {
                $user->setFirstName($firstName);
            }

            if (!empty($lastName)) {
                $user->setLastName($lastName);
            }

            if (!empty($email)) {
                $user->setEmail($email);
            }

            if (!empty($phone)) {
                $user->setPhone($phone);
            }

            $this->userRepository->update($user);

            return $this->json([
                'id' => $user->getId()
            ]);
        } catch (\Throwable $e) {
            $errorsCounter = $this->getApiMethodMetricsCounter(__FUNCTION__, 'errors');
            $errorsCounter->inc();

            return $this->json([
                'code' => 500,
                'message' => $e->getMessage()
            ]);
        }
    }

    /**
     * @Route("/user/{id<\d+>}", name="deleteUserAction", methods={"DELETE"})
     *
     * @param int $id
     *
     * @return JsonResponse
     */
    public function deleteUserAction(int $id): JsonResponse
    {
        $rpsCounter = $this->getApiMethodMetricsCounter(__FUNCTION__, 'rps');
        $rpsCounter->inc();

        try {
            $user = $this->userRepository->findById($id);
            if ($user === null) {
                return $this->json([
                    'code' => 0,
                    'message' => 'User not found'
                ]);
            }

            $this->userRepository->delete($user);
            return $this->json([
                'code' => 0,
                'message' => 'Success!'
            ]);
        } catch (\Throwable $e) {
            $errorsCounter = $this->getApiMethodMetricsCounter(__FUNCTION__, 'errors');
            $errorsCounter->inc();

            return $this->json([
                'code' => 500,
                'message' => $e->getMessage()
            ]);
        }
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
     * @param string $apiMethod
     *
     * @return Counter
     *
     * @throws \Prometheus\Exception\MetricsRegistrationException
     */
    private function getApiMethodMetricsCounter(string $apiMethod, string $event): Counter
    {
        return $this->prometheusRegistry->registerCounter(
            $this->metricsMap['namespace'],
            $this->metricsMap[$apiMethod][$event]['name'],
            $this->metricsMap[$apiMethod][$event]['help'],
            $this->metricsMap[$apiMethod][$event]['labels']
        );
    }
}