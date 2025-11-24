#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

// Configuration
const HARDCOVER_API_URL = 'https://api.hardcover.app/v1/graphql';
const ABOUT_PAGE_PATH = path.join(__dirname, 'content/about/_index.md');

// Get API token from environment variable
const API_TOKEN = process.env.HARDCOVER_API_TOKEN;

if (!API_TOKEN) {
  console.error('Error: HARDCOVER_API_TOKEN environment variable is not set');
  console.error('Please set it with: export HARDCOVER_API_TOKEN="your-token-here"');
  process.exit(1);
}

// GraphQL query to get currently reading books
const query = `
  query {
    me {
      user_books(where: {status_id: {_eq: 2}}) {
        book {
          title
          cached_image
          image {
            url
            width
            height
          }
          contributions {
            author {
              name
            }
          }
        }
      }
    }
  }
`;

async function fetchCurrentlyReading() {
  try {
    const response = await fetch(HARDCOVER_API_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${API_TOKEN}`,
      },
      body: JSON.stringify({ query }),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();

    if (data.errors) {
      throw new Error(`GraphQL errors: ${JSON.stringify(data.errors)}`);
    }

    return data.data.me[0].user_books;
  } catch (error) {
    console.error('Error fetching currently reading books:', error);
    throw error;
  }
}

function updateAboutPage(bookInfo) {
  try {
    // Read the current about page
    const content = fs.readFileSync(ABOUT_PAGE_PATH, 'utf8');

    // Split frontmatter and body
    const parts = content.split('---');
    if (parts.length < 3) {
      throw new Error('Invalid frontmatter format');
    }

    const frontmatter = parts[1];
    const body = parts.slice(2).join('---');

    // Parse and update frontmatter
    const lines = frontmatter.split('\n');
    let updatedLines = [];
    let inCurrentlyReading = false;
    let currentlyReadingIndent = 0;

    for (let i = 0; i < lines.length; i++) {
      const line = lines[i];

      if (line.trim().startsWith('currentlyReading:')) {
        inCurrentlyReading = true;
        currentlyReadingIndent = line.search(/\S/);
        updatedLines.push(line);

        // Add updated book info
        updatedLines.push(`  title: "${bookInfo.title}"`);
        updatedLines.push(`  author: "${bookInfo.author}"`);
        updatedLines.push(`  coverImage: "${bookInfo.coverImage}"`);

        // Skip the old currentlyReading fields
        let j = i + 1;
        while (j < lines.length && lines[j].trim() && !lines[j].match(/^\w/)) {
          j++;
        }
        i = j - 1;
      } else if (!inCurrentlyReading) {
        updatedLines.push(line);
      } else {
        if (line.match(/^\w/)) {
          inCurrentlyReading = false;
          updatedLines.push(line);
        }
      }
    }

    // Reconstruct the file
    const updatedContent = `---${updatedLines.join('\n')}---${body}`;

    // Write back to file
    fs.writeFileSync(ABOUT_PAGE_PATH, updatedContent, 'utf8');
    console.log('✓ Successfully updated about page with currently reading book');
  } catch (error) {
    console.error('Error updating about page:', error);
    throw error;
  }
}

async function main() {
  try {
    console.log('Fetching currently reading books from Hardcover...');
    const userBooks = await fetchCurrentlyReading();

    if (!userBooks || userBooks.length === 0) {
      console.log('No books currently reading. Leaving about page unchanged.');
      return;
    }

    // Get the first book (0th index)
    const firstBook = userBooks[0].book;

    // Extract author - use first one (0th index)
    let author = 'Unknown Author';
    if (firstBook.contributions && firstBook.contributions.length > 0) {
      author = firstBook.contributions[0].author.name;
    }

    // Get cover image URL
    const coverImage = firstBook.image?.url || '';

    const bookInfo = {
      title: firstBook.title,
      author: author,
      coverImage: coverImage,
    };

    console.log(`Found currently reading: "${bookInfo.title}" by ${bookInfo.author}`);

    // Update the about page
    updateAboutPage(bookInfo);

  } catch (error) {
    console.error('Failed to update currently reading book:', error);
    process.exit(1);
  }
}

main();
