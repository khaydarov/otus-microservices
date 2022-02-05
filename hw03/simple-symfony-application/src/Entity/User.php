<?php

declare(strict_types=1);

namespace App\Entity;

use InvalidArgumentException;

class User
{
    /**
     * @var int
     */
    private int $id;

    /**
     * @var string
     */
    private string $username;

    /**
     * @var string
     */
    private string $firstName;

    /**
     * @var string
     */
    private string $lastName;

    /**
     * @var string
     */
    private string $email;

    /**
     * @var string
     */
    private string $phone;

    public function __construct(
        int $id,
        string $username,
        string $firstName,
        string $lastName,
        string $email,
        string $phone
    ) {
        $this->id = $id;

        if (empty($username)) {
            throw new InvalidArgumentException("Username is empty");
        }
        $this->username = $username;

        if (empty($firstName)) {
            throw new InvalidArgumentException("First name is empty");
        }
        $this->firstName = $firstName;

        if (empty($lastName)) {
            throw new InvalidArgumentException("Last name is empty");
        }
        $this->lastName = $lastName;

        if (empty($email)) {
            throw new InvalidArgumentException("Email is empty");
        }
        $this->email = $email;

        if (empty($phone)) {
            throw new InvalidArgumentException("Phone is empty");
        }
        $this->phone = $phone;
    }

    /**
     * @return int
     */
    public function getId(): int
    {
        return $this->id;
    }

    /**
     * @return string
     */
    public function getUsername(): string
    {
        return $this->username;
    }

    /**
     * @param string $username
     * @return $this
     */
    public function setUsername(string $username): self
    {
        $this->username = $username;

        return $this;
    }

    /**
     * @return string
     */
    public function getFirstName(): string
    {
        return $this->firstName;
    }

    /**
     * @param string $firstname
     * @return $this
     */
    public function setFirstName(string $firstname): self
    {
        $this->firstName = $firstname;

        return $this;
    }

    /**
     * @return string
     */
    public function getLastName(): string
    {
        return $this->lastName;
    }

    /**
     * @param string $lastName
     * @return $this
     */
    public function setLastName(string $lastName): self
    {
        $this->lastName = $lastName;

        return $this;
    }

    /**
     * @return string
     */
    public function getEmail(): string
    {
        return $this->email;
    }

    /**
     * @param string $email
     * @return $this
     */
    public function setEmail(string $email): self
    {
        $this->email = $email;

        return $this;
    }

    /**
     * @return string
     */
    public function getPhone(): string
    {
        return $this->phone;
    }

    /**
     * @param string $phone
     * @return $this
     */
    public function setPhone(string $phone): self
    {
        $this->phone = $phone;

        return $this;
    }
}