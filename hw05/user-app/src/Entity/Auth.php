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
    private $email;

    public function __construct(int $id, string $email)
    {
        $this->id = $id;
        $this->email = $email;
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
    public function getEmail(): string
    {
        return $this->email;
    }
}