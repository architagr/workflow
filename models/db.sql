
create database workflow
go
use workflow
go

create table dbo.Licenses
(
   Id int identity(1,1) primary key,
   LicenseName nvarchar(200) not null,
   IsActive bit default(1) not null,
   EffectiveValidDate date default(getdate()) not null,
   DiscontinueDate date null,
   Cost int null,
   LoginCount int null,
   OficeCount int null,
   DefaultDuration int NULL,
   ChildFirmCount INT NULL,
	IsTrial BIT NOT NULL DEFAULT(0)
);

create table dbo.Countries
(
   Id int identity(1,1) primary key,
   Name nvarchar(200) not null,
   AliasName nvarchar(200) not null,
   IsActive bit default(1) not null
);

create table dbo.States
(
   Id int identity(1,1) primary key,
   Name nvarchar(200) not null,
   AliasName nvarchar(200) not null,
   CountryId int references dbo.Countries(Id) not null,
   IsActive bit default(1) not null
);

create table dbo.Cities
(
   Id int identity(1,1) primary key,
   Name nvarchar(200) not null,
   AliasName nvarchar(200) not null,
   CountryId int references dbo.Countries(Id) not null,
   StateId int references dbo.States(Id) not null,
   DefaultPostalCode nvarchar(200) not null,
   IsCapital bit default(0) not null,
   IsActive bit default(1) not null
);

create table dbo.RegistrationAgencies
(
   Id int identity(1,1) primary key,
   Name nvarchar(200) not null,
   AliasName nvarchar(200) not null,
   CountryId int references dbo.Countries(Id) not null,
   StateId int references dbo.States(Id) not null,
   CityId int references dbo.Cities(Id) not null,
   IsActive bit default(1) not null
);

create table dbo.Tenant
(
   Id int identity(1,1) primary key,
   CompanyName nvarchar(500) not null,
   AliasName nvarchar(200) not null,
   RegistrationNumber nvarchar(500) not null,
   LogoPathLarge nvarchar(1000),
   LogoPathSmall nvarchar(1000),
   RegistrationAgency int references RegistrationAgencies(Id) null,
   RegisteredEmail nvarchar(200),
   DBConnectionString nvarchar(1000),
	DBType INT,
   AccountStatus tinyInt,
   IsActive bit default(1),
   CreateDate datetime default(getdate()) not null,
   CreatedBy int not null,
   UpdateDate datetime null,
   UpdateBy int not null
);

create table dbo.Users
(
   Id int identity(1,1) primary key,
   TenantId int references dbo.Tenant(Id) null,
   Name nvarchar(500) not null,
   PrimaryEmail nvarchar(200),
   SecondaryEmail nvarchar(200),
   PrimaryMobile nvarchar(20),
   SecondaryMobile nvarchar(20),
   FixedLine nvarchar(20),
   FixedLineExt nvarchar(20),
   [Role] tinyint default(4) not null,
   LoginId nvarchar(200) not null,
   [Password] nvarchar(500) not null,
   PasswordSetDate date null,
   IsFirstLogin bit default(1) not null,
   FirstLoginDate datetime null,
   IsResetPassword bit default(0) not null,
   ResetPasswordRequestDate datetime null,
   ResetPasswordCode uniqueidentifier null,
   IsActive bit default(1),
   CreateDate datetime default(getdate()) not null,
   CreatedBy int null,
   UpdateDate datetime null,
   UpdateBy int null
);

insert into dbo.Users
values(null, 'Super Admin', 'superadmin@gmail.com', 'superadmin@gmail.com', '8939354552', '8939354552',
       '8939354552', '8939354552', 1, 'admin', 'adminP', null, 1, null, 1, null, NEWID(), 1, getdate(), null, null, null);

ALTER TABLE dbo.Users ADD FOREIGN KEY (CreatedBy) REFERENCES Users(Id);

ALTER TABLE dbo.Users ADD FOREIGN KEY (UpdateBy) REFERENCES Users(Id);

ALTER TABLE dbo.Tenant ADD FOREIGN KEY (CreatedBy) REFERENCES Users(Id);

ALTER TABLE dbo.Tenant ADD FOREIGN KEY (UpdateBy) REFERENCES Users(Id);

create table dbo.TenantLicense
(
   Id int identity(1,1) primary key,
   TenantId int references dbo.Tenant(Id) not null,
   LicenseId int references dbo.Licenses(Id) not null,
   StartDate date default(getdate()) not null,
   EndtDate date null,
   EndReason nvarchar(1000),
   Notes nvarchar(max),
   Discount int default(0) not null,
   CreateDate datetime default(getdate()) not null,
   UpdateDate datetime null,
   UpdateBy int references dbo.Users(Id) not null
);

create table dbo.Addresses
(
   Id int identity(1,1) primary key,
   TenantId int references dbo.Tenant(Id) not null,
   UserId int references dbo.Users(Id) null,
   AddressLine1 nvarchar(500),
   AddressLine2 nvarchar(500),
   AddressLine3 nvarchar(500),
   CountryId int references dbo.Countries(Id) not null,
   StateId int references dbo.States(Id) not null,
   CityId int references dbo.Cities(Id) not null,
   IsActive bit default(1) not null,
   UpdateDate datetime null,
   UpdateBy int references dbo.Users(Id) not null
);

create table dbo.OrganizationHierarchy
(
   Id int identity(1,1) primary key,
	[Name] nvarchar(500) not null,
   ParentId int null,
   HasChild bit default(0),
   TenantId int references dbo.Tenant(Id) not null,
	IsActive bit default(1) not null,
	CreateDate datetime default(getdate()) not null,
	CreatedBy int references dbo.Users(Id) not null,
    UpdateDate datetime null,
    UpdateBy int references dbo.Users(Id) null
)

create table dbo.UserOrganizationHierarchy
(
   Id int identity(1,1) primary key,
	UserId int references dbo.Users(Id) not null,
OrganizationHierarchyId int references dbo.OrganizationHierarchy(Id) not null,
	CreateDate datetime default(getdate()) not null,
	CreatedBy int references dbo.Users(Id) not null,
    UpdateDate datetime null,
    UpdateBy int references dbo.Users(Id) null
)
