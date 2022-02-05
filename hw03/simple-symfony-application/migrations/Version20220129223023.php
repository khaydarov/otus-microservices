<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20220129223023 extends AbstractMigration
{
    public function getDescription(): string
    {
        return 'Creates Users table';
    }

    public function up(Schema $schema): void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql('
            CREATE TABLE IF NOT EXISTS t_users (
                id bigserial NOT NULL PRIMARY KEY,
                username varchar(48) NOT NULL,
                firstName varchar(48) NOT NULL,
                lastName varchar(48) NOT NULL,
                email text NOT NULL,
                phone text NOT NULL
            )
        ');
    }

    public function down(Schema $schema): void
    {
        // this down() migration is auto-generated, please modify it to your needs
    }
}
