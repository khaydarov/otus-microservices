<?php

declare(strict_types=1);

namespace App\Repository;

use App\Entity\User;
use Doctrine\DBAL\Connection;

class UserRepository
{
    /**
     * @var Connection
     */
    private Connection $dbal;

    public function __construct(Connection $dbal)
    {
        $this->dbal = $dbal;
    }

    /**
     * @param int $id
     *
     * @return User|null
     *
     * @throws \Doctrine\DBAL\Exception
     */
    public function findById(int $id): ?User
    {
        $row = $this->dbal->fetchAssociative("SELECT * FROM t_users WHERE id = :id", [
            'id' => $id
        ]);

        if (empty($row)) {
            return null;
        }

        return new User(
            $row['id'],
            $row['username'],
            $row['firstname'],
            $row['lastname'],
            $row['email'],
            $row['phone']
        );
    }

    /**
     * @param User $user
     *
     * @throws \Doctrine\DBAL\Exception
     */
    public function insert(User $user): void
    {
        $this->dbal->executeQuery(
            "INSERT INTO t_users (id, username, firstname, lastname, email, phone)
            VALUES (:id, :username, :firstname, :lastname, :email, :phone)
        ",
        [
            'id' => $user->getId(),
            'username' => $user->getUsername(),
            'firstname' => $user->getFirstName(),
            'lastname' => $user->getLastName(),
            'email' => $user->getEmail(),
            'phone' => $user->getPhone(),
        ]);
    }

    /**
     * @param User $user
     *
     * @throws \Doctrine\DBAL\Exception
     */
    public function update(User $user): void
    {
        $this->dbal->executeQuery(
        "UPDATE t_users SET 
                username = :username,
                firstname = :firstname,
                lastname = :lastname,
                email = :email,
                phone = :phone
                WHERE id = :id
            ",
            [
                'id' => $user->getId(),
                'username' => $user->getUsername(),
                'firstname' => $user->getFirstName(),
                'lastname' => $user->getLastName(),
                'email' => $user->getEmail(),
                'phone' => $user->getPhone()
            ]
        );
    }

    /**
     * @param User $user
     */
    public function delete(User $user): void
    {
        $this->dbal->executeQuery("DELETE FROM t_users WHERE id = :id", [
            'id' => $user->getId()
        ]);
    }


    /**
     * @return int
     *
     * @throws \Doctrine\DBAL\Exception
     */
    public function nextIdentity(): int
    {
        return $this->dbal->fetchOne("select nextval('t_users_id_seq'::regclass)");
    }
}