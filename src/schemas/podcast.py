"""
Podcasts schemas
"""

from xml.etree.ElementTree import Element, SubElement

from pydantic import BaseModel


class PodcastOut(BaseModel):
    """Podcast schema"""

    name: str
    description: str
    website: str
    xml_url: str
    author: str
    subscribers_count: int

    def to_opml(self) -> str:
        """Convert Podcast to OPML outline element"""
        outline = Element(
            "outline",
            text=self.name,
            description=self.description,
            xmlUrl=self.xml_url,
            htmlUrl=self.website,
        )
        return outline

    def to_xml(self) -> Element:
        """Convert Podcast to XML format"""
        podcast = Element("podcast")
        name = SubElement(podcast, "name")
        name.text = self.name
        description = SubElement(podcast, "description")
        description.text = self.description
        website = SubElement(podcast, "website")
        website.text = self.website
        xml_url = SubElement(podcast, "xml_url")
        xml_url.text = self.xml_url
        author = SubElement(podcast, "author")
        author.text = self.author

        return podcast
