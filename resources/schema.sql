CREATE TABLE person (
    id integer NOT NULL,
    firstname varchar(255) NOT NULL,
    lastname varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    priv integer DEFAULT 0,
    created timestamp DEFAULT now() NOT NULL,
    updated timestamp DEFAULT now() NOT NULL,
    primary key (id)
);

CREATE TABLE account (
    id integer NOT NULL,
    account varchar(255) NOT NULL,
    person_id integer NOT NULL,
    created timestamp DEFAULT now() NOT NULL,
    updated timestamp DEFAULT now() NOT NULL,
    foreign key (person_id) references person(id),
    primary key (id)
);

CREATE TABLE campaign (
    id integer NOT NULL,
    campaign varchar(255) NOT NULL,
    account_id integer NOT NULL,
    created timestamp DEFAULT now() NOT NULL,
    updated timestamp DEFAULT now() NOT NULL,
    foreign key (account_id) references account(id),
    primary key (id)
);

CREATE TABLE survey (
    id integer NOT NULL,
    survey varchar(255) NOT NULL,
    account_id integer NOT NULL,
    questions json,
    -- questions json -- built from question
    status integer NOT NULL DEFAULT 0,
    created timestamp DEFAULT now() NOT NULL,
    updated timestamp DEFAULT now() NOT NULL,
    foreign key (account_id) references account(id),
    primary key (id)
);

CREATE TABLE survey_question ( -- many to many
    survey_id integer REFERENCES survey(id),
    question_id integer REFERENCES question(id),
    PRIMARY KEY (survey_id, question_id)
);

CREATE TABLE question (
    id integer NOT NULL,
    created timestamp DEFAULT now() NOT NULL,
    updated timestamp DEFAULT now() NOT NULL,
    primary key (id)
);

CREATE TABLE answer (
    id integer NOT NULL,
    survey_id integer NOT NULL,
    question_id integer NOT NULL,
    created timestamp DEFAULT now() NOT NULL,
    updated timestamp DEFAULT now() NOT NULL,
    foreign key (question_id) references question(id),
    primary key (id)
);

-- Data for testing --
