<?php

declare(strict_types=1);

namespace App\Controller;

use App\Entity\User;
use App\Repository\UserRepository;
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
    public function getUserAction(int $id): JsonResponse
    {
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
}