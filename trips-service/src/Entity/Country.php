<?php

namespace App\Entiy;

use DateTimeImmutable;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping\Entity;
use Doctrine\ORM\Mapping\Table;
use Doctrine\ORM\Mapping;
use Doctrine\DBAL\Types\Types;
use Gedmo\Mapping\Annotation as Gedmo;

#[Entity]
#[Table(name: "country")]
class Country
{
    public function __construct()
    {
        $this->places = new ArrayCollection();
    }

    #[Mapping\Id]
    #[Mapping\GeneratedValue(strategy: "AUTO")]
    #[Mapping\Column(type: Types::INTEGER, name: "country_id" )]
    private ?int $countryId = null;

    #[Mapping\Column(type: Types::TEXT, unique: true)]
    private ?string $code = null;

    #[Mapping\Column(type: Types::TEXT, unique: true)]
    private ?string $name = null;

    #[Mapping\ManyToOne(inversedBy: "place")]
    private Collection $places;

    #[Gedmo\Timestampable(on: 'create')]
    #[Mapping\Column(name: "created_at", type: Types::DATETIME_IMMUTABLE)]
    private ?DateTimeImmutable $createdAt = null;

    #[Gedmo\Timestampable(on: 'update')]
    #[Mapping\Column(name: "updated_at", type: Types::DATETIME_IMMUTABLE, nullable: true)]
    private ?DateTimeImmutable $updatedAt = null;

    #[Mapping\Column(name: "deleted_at", type: Types::DATETIME_IMMUTABLE, nullable: true)]
    private ?DateTimeImmutable $deletedAt = null;

    public function getCountryId(): ?int 
    {
        return $this->countryId;
    }

    public function getCode(): ?string
    {
        return $this->code;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    /**  @return Collection<int, Place>*/
    public function getPlaces(): Collection
    {
        return $this->places;
    }

    public function addPlace(Place $place): self
    {
        if ($this->places->contains($place)) {
            return $this;
        }

        $this->places->add($place);
        $place->setCountry($this);

        return $this;
    }

    public function setDeletedAt(DateTimeImmutable $date): self
    {
        $this->deletedAt = $date;

        return $this;
    }

    public function removePlace(Place $place): self
    {
        if (!$this->places->contains($place)) {
            return $this;
        }

        $this->places->removeElement($place);
        $place->setCountry(null);

        return $this;
    }

    public function getCreatedAt(): ?DateTimeImmutable
    {
        return $this->createdAt;
    }

    public function getUpdatedAt(): ?DateTimeImmutable
    {
        return $this->updatedAt;
    }

    public function getDeletedAt(): ?DateTimeImmutable
    {
        return $this->deletedAt;
    }
}