<?php 

namespace App\Entity;

use DateTimeImmutable;
use Doctrine\ORM\Mapping\Entity;
use Doctrine\ORM\Mapping\Table;
use Doctrine\ORM\Mapping;
use Doctrine\DBAL\Types\Types;
use Gedmo\Mapping\Annotation as Gedmo;

#[Entity]
#[Table(name: "trip")]
class Trip
{
    #[Mapping\Id]
    #[Mapping\Column(name: "trip_id")]
    #[Mapping\GeneratedValue(strategy: "AUTO")]
    private ?int $tripId = null;

    #[Mapping\Column(type: Types::TEXT, unique: true)]
    private ?string $name = null;

    #[Mapping\Column(type: Types::INTEGER, nullable: false)]
    private ?int $authorId = null;

    #[Gedmo\Timestampable(on: 'create')]
    #[Mapping\Column(name: "created_at", type: Types::DATETIME_IMMUTABLE)]
    private ?DateTimeImmutable $createdAt = null;

    #[Gedmo\Timestampable(on: 'update')]
    #[Mapping\Column(name: "updated_at", type: Types::DATETIME_IMMUTABLE, nullable: true)]
    private ?DateTimeImmutable $updatedAt = null;

    #[Mapping\Column(name: "deleted_at", type: Types::DATETIME_IMMUTABLE, nullable: true)]
    private ?DateTimeImmutable $deletedAt = null;

    #[Mapping\Column(name: "expirse_at", type: Types::DATETIME_IMMUTABLE, nullable: true)]
    private ?DateTimeImmutable $expirseAt = null;

    #[Mapping\OneToMany(mappedBy: 'place', targetEntity: Place::class)]
    #[Mapping\JoinColumn(nullable: false)]
    private $place;

    public function getTripId(): ?int
    {
        return $this->tripId;
    }

    public function getCreatedAt(): ?DateTimeImmutable
    {
        return $this->createdAt;
    }

    public function getExpirseAt(): ?DateTimeImmutable
    {
        return $this->expirseAt;
    }

    public function getUpdatedAt(): ?DateTimeImmutable
    {
        return $this->updatedAt;
    }

    public function getDeletedAt(): ?DateTimeImmutable
    {
        return $this->deletedAt;
    }

    public function getPlace(): ?Place
    {
        return $this->place;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    public function setName(string $name): self
    {
        $this->name = $name;

        return $this;
    }

    public function getAuthorId(): ?int
    {
        return $this->authorId;
    }

    public function setAuthorId(int $authorId): self
    {
        $this->authorId = $authorId;

        return $this;
    }
}