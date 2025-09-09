<?php 

namespace App\Entity;

use DateTimeImmutable;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping\Entity;
use Doctrine\ORM\Mapping\Table;
use Doctrine\ORM\Mapping;
use Doctrine\DBAL\Types\Types;
use Gedmo\Mapping\Annotation as Gedmo;

#[Entity]
#[Table(name: "place")]
class Place
{
    #[Mapping\Id]
    #[Mapping\Column(name: "place_id")]
    #[Mapping\GeneratedValue(strategy: "AUTO")]
    private ?int $placeId = null;

    #[Gedmo\Timestampable(on: 'create')]
    #[Mapping\Column(name: "created_at", type: Types::DATETIME_IMMUTABLE)]
    private ?DateTimeImmutable $createdAt = null;

    #[Gedmo\Timestampable(on: 'update')]
    #[Mapping\Column(name: "updated_at", type: Types::DATETIME_IMMUTABLE, nullable: true)]
    private ?DateTimeImmutable $updatedAt = null;

    #[Mapping\Column(name: "deleted_at", type: Types::DATETIME_IMMUTABLE, nullable: true)]
    private ?DateTimeImmutable $deletedAt = null;

    #[Mapping\Column(name: "code", type: Types::TEXT)]
    private ?string $code = null;

    #[Mapping\Column(name: "name", type: Types::TEXT)]
    private ?string $name = null;

    #[Mapping\OneToMany(mappedBy: 'country', targetEntity: Country::class)]
    #[Mapping\JoinColumn(nullable: false)]
    private ?Country $country = null;

    /** @return Collection<int, Trip> */
    private Collection $trips;

    public function __construct()
    {
        $this->trips = new Collection();
    }

    public function getPlaceId(): ?int
    {
        return $this->placeId;
    }

    public function getCode(): ?string
    {
        return $this->code;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    public function getCreatedAt(): ?DateTimeImmutable
    {
        return $this->createdAt;
    }

    public function getCountry(): ?Country
    {
        return $this->country;
    }

    public function getUpdatedAt(): ?DateTimeImmutable
    {
        return $this->updatedAt;
    }

    public function getDeletedAt(): ?DateTimeImmutable
    {
        return $this->deletedAt;
    }

    public function getTrips(): Collection
    {
        return $this->trips;
    }

    public function addTrip(Trip $trip): self
    {
        if ($this->trips->contains($trip)) {
            return $this;
        }

        $this->trips->add($trip);

        return $this;
    }

    public function setDeletedAt(DateTimeImmutable $date): self
    {
        $this->deletedAt = $date;

        return $this;
    }

    public function setCountry(?Country $country): self
    {
        $this->country = $country;

        return $this;
    }
}