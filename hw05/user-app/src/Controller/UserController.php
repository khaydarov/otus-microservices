<?php

declare(strict_types=1);

namespace App\Controller;

use App\Entity\Auth;
use App\Entity\User;
use App\Repository\UserRepository;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;

/**
 * Class UserController
 *
 * @package App\Controller
 */
final class UserController extends AbstractController
{
    /**
     * @var UserRepository
     */
    private UserRepository $userRepository;

    public function __construct(UserRepository $userRepository)
    {
        $this->userRepository = $userRepository;
    }

    /**
     * @Route("/users", name="getUsersAction", methods={"GET"})
     *
     * @param Request $request
     *
     * @return JsonResponse
     */
    public function getUsersAction(Request $request): JsonResponse
    {
        $email = $request->query->get('email', '');
        $password = $request->query->get('password', '');

        $user = $this->userRepository->findByEmailAndPassword($email, $password);
        if ($user === null) {
            return $this->json([
                'code' => 0,
                'message' => 'User not found'
            ]);
        }

        return $this->json([
            'id' => $user->getId(),
            'email' => $user->getEmail(),
        ]);
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
        $requestData = json_decode($request->getContent(), true);

        $username = $requestData['username'] ?? '';
        $firstName = $requestData['firstName'] ?? '';
        $lastName = $requestData['lastName'] ?? '';
        $email = $requestData['email'] ?? '';
        $phone = $requestData['phone'] ?? '';

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
    }

    /**
     * @Route("/user/{id<\d+>}", name="getUserAction", methods={"GET"})
     *
     * @param int $id
     *
     * @return JsonResponse
     */
    public function getUserAction(Request $request, int $id): JsonResponse
    {
        $auth = $this->getAuthData($request);
        if ($auth === null || $auth->getId() !== $id) {
            return $this->json([
                'code' => 403,
                'message' => 'Access denied'
            ]);
        }

        $user = $this->userRepository->findById($id);

        if ($user === null) {
            return $this->json([
                'code' => 404,
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
    }

    /**
     * @Route("/user/{id<\d+>}", name="putUserAction", methods={"PUT"})
     *
     * @param Request $request
     * @param int $id
     *
     * @return JsonResponse
     */
    public function putUserAction(Request $request, int $id): JsonResponse
    {
        try {
            $auth = $this->getAuthData($request);
            if ($auth === null || $auth->getId() !== $id) {
                return $this->json([
                    'code' => 403,
                    'message' => 'Access denied'
                ]);
            }

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
            return $this->json([
                'code' => 500,
                'message' => $e->getMessage()
            ]);
        }
    }

    /**
     * @param Request $request
     * @return Auth|null
     */
    private function getAuthData(Request $request): ?Auth
    {
        $authToken = $request->headers->get('x-auth-token', '');
        if (empty($authToken)) {
            return null;
        }

        $decoded = (array) JWT::decode($authToken, new Key($this->getParameter('auth_salt'), 'HS256'));
        $now = new \DateTimeImmutable();
        $expiration = new \DateTimeImmutable();
        $expiration->setTimestamp($decoded['expiration_in']);

        if ($now > $expiration) {
            throw new \InvalidArgumentException('Token is expired');
        }

        return new Auth($decoded['user_id'], $decoded['user_email']);
    }
}