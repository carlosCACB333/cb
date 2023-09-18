import { SkillSection } from "@/components/home/skill-section";
import { Hero } from "@/components/home/hero";
import { AboutSection } from "@/components/home/hero/about-section";
import { env } from "@/utils";
import {
  AiFillSafetyCertificate,
  AiOutlineFundProjectionScreen,
} from "react-icons/ai";
import { BsFillPostcardFill } from "react-icons/bs";
import { ProjectSection } from "@/components/home/project-section";
import { CertificateSection } from "@/components/home/certificates-section";
import { ContactSection } from "@/components/home/contact-section";
import { Locale, Stage } from "@/generated/graphql";
import { getSdk } from "@/utils/sdk";

export default async function Home() {
  const {
    categories,
    certifications,
    certificationsConnection,
    postsConnection,
    projectsConnection,
    projects,
  } = await getSdk().getHomeData(
    {
      locales: [Locale.Es],
      stage: Stage.Published,
    },
    {}
  );

  return (
    <main className="container mx-auto max-w-7xl px-6 flex-grow">
      <Hero
        features={[
          {
            icon: <AiFillSafetyCertificate size={32} />,
            title: "Certificaciones",
            description: certificationsConnection?.aggregate?.count + "+",
            href: "#home-certifications",
          },
          {
            icon: <AiOutlineFundProjectionScreen size={32} />,
            title: "Proyectos",
            description: projectsConnection?.aggregate?.count + "+",
            href: "#home-projects",
          },
          {
            icon: <BsFillPostcardFill size={32} />,
            title: "Publicaciones",
            description: postsConnection?.aggregate?.count + "+",
            href: "/blog",
          },
        ]}
      />
      <AboutSection />
      <SkillSection categories={categories as any} />
      <ProjectSection projects={projects as any} />
      <CertificateSection certifications={certifications as any} />
      <ContactSection />
    </main>
  );
}

export const revalidate = env.revalidate;
