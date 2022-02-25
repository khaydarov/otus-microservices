<?php

declare(strict_types=1);

namespace App\Entity;

/**
 * Class Auth
 * @package App\Entity
 */
final class Auth
{
    /**
     * @var int
     */
    private $id;

    /**
     * @var string
     */
    private $username;

    public function __construct(int $id, string $username)
    {
        $this->id = $id;
        $this->username = $username;
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
}